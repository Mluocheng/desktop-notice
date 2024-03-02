import './App.css';
import Header from './components/header';
import Body from './components/body';
import Footer from './components/footer';
import { useHover } from 'ahooks';
import { useCallback, useEffect, useState } from 'react';
import { GetWindow, HideWindow, ShowWindow } from '../wailsjs/go/window/Window';
import { request } from './utils/request';

let closeWindowTimer: number | null | undefined = null;

function App() {
    const [wailsData, setWailsData] = useState<WailsProps>({})
    const isHovering = useHover(() => document.getElementById('root'));

    useEffect(() => {
        // ShowWindow();
        initWindowUtils();
    }, [])

    function initWindowUtils() {
        window.DesktopNotice = {
            myGlobalFunction: () => console.log('myGlobalFunction'),
            closeWindow: () => window.runtime.Quit(),
            request: (url, options) => request(url, options),
        };
    }

    // 默认3秒关闭
    useEffect(() => {
        getWailData();
        return () => {
            clearCloseWindowTimer();
        };
    }, [])

    // const initPosition = useCallback(() => {
    //     const x = (screen.width - (wailsData?.Width || 336)) * window.devicePixelRatio - 16
    //     const y = (screen.height - (wailsData?.Height || 200)) * window.devicePixelRatio - 76
    //     window.runtime.WindowSetPosition(x, y)
    // }, [wailsData?.Width, wailsData?.Height])

    useEffect(() => {
        console.log("定时器", wailsData.AutoCloseWindowTimer, wailsData.AutoCloseWindowTimer !== -1)
        if (wailsData.AutoCloseWindowTimer && (wailsData.AutoCloseWindowTimer !== -1)) {
            if (isHovering) {
                console.log('isHovering', isHovering, closeWindowTimer, "清除倒计时关闭")
                clearCloseWindowTimer()
            } else {
                startCloseWindowTime()
            }
        }
    }, [isHovering, wailsData.AutoCloseWindowTimer])


    function startCloseWindowTime() {
        closeWindowTimer = setTimeout(() => {
            console.log('关闭窗口', closeWindowTimer)
            if (!isHovering) {
                // @ts-ignore
                if (process.env.NODE_ENV === "development") {
                    // HideWindow() // todo* 测试关闭先
                } else {
                    window.runtime.Quit()
                }
            }
        }, wailsData?.AutoCloseWindowTimer || 5000)
    }

    function clearCloseWindowTimer() {
        if (closeWindowTimer) {
            if (closeWindowTimer) clearTimeout(closeWindowTimer);
            closeWindowTimer = null;
        }
    }

    // 手动关闭
    function handelClose() {
        // @ts-ignore
        if (process.env.NODE_ENV === "development") {
            HideWindow()
        } else {
            window.runtime.Quit()
        }
        clearCloseWindowTimer()
    }

    async function getWailData() {
        const data = await GetWindow()
        console.log("data.Data", data.Data)
        if (data.Data) setWailsData(data.Data)

        const x = (screen.width - (data.Data?.Width || 336)) * window.devicePixelRatio - 16
        const y = (screen.height - (data.Data?.Height || 200)) * window.devicePixelRatio - 76
        window.runtime.WindowSetPosition(x, y)
    }


    return (
        <div id="App" className='app'>
            <Header {...wailsData} handelClose={handelClose} />
            <Body {...wailsData} />
            <Footer />
        </div>
    )
}

export default App;
