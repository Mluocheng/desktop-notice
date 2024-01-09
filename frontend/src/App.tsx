import './App.css';
import Header from './components/header';
import Body from './components/body';
import Footer from './components/footer';
import { useHover } from 'ahooks';
import { useEffect } from 'react';
import { HideWindow } from '../wailsjs/go/window/Window';

let closeWindow: number | null | undefined = null;
function App() {
    
    const isHovering = useHover(() => document.getElementById('root'), {
        onEnter: () => {
            console.log('onEnter');
        },
        onLeave: () => {
            console.log('onLeave');
        },
        onChange: isHover => {
            console.log('onChange', isHover);
        },
    });
    console.log(isHovering)

    // 默认3秒关闭
    useEffect(() => {
        handleCloseWindow()
    }, [])

    useEffect(() => {
        if (isHovering) {
            if (closeWindow) clearTimeout(closeWindow)
        } else {
            handleCloseWindow()
        }
    }, [isHovering])

    // 关闭窗口函数
    function handleCloseWindow() {
        closeWindow = setTimeout(() => {
            HideWindow()
        }, 3000)
    }

    return (
        <div id="App" className='app'>
            <Header />
            <Body />
            <Footer />
        </div>
    )
}

export default App;
