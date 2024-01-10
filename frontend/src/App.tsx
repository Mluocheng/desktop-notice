import './App.css';
import Header from './components/header';
import Body from './components/body';
import Footer from './components/footer';
import { useHover } from 'ahooks';
import { useCallback, useEffect, useState } from 'react';
import { GetWindow, HideWindow } from '../wailsjs/go/window/Window';

let closeWindowTimer: number | null | undefined = null;

function App() {
    const [wailsData, setWailsData] = useState<WailsProps>({})
    const isHovering = useHover(() => document.getElementById('root'));

    // 默认3秒关闭
    useEffect(() => {
        initPosition()
        getWailData();
        return () => {
            clearCloseWindowTimer();
        };
    }, [])

    function initPosition() {
        const x = (screen.width - 336) * window.devicePixelRatio - 16
        const y = (screen.height - 200) * window.devicePixelRatio - 76
        window.runtime.WindowSetPosition(x, y)
    }

    useEffect(() => {
        if (isHovering) {
            console.log('isHovering', isHovering, closeWindowTimer, "清除倒计时关闭")
            clearCloseWindowTimer()
        } else {
            startCloseWindowTime()
        }
    }, [isHovering])


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
        }, 5000)
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
        if (data.Data) setWailsData(data.Data)
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
