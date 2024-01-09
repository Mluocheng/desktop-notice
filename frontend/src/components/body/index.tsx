import { useEffect } from 'react';
import './index.css';
import { ShowWindow, HideWindow } from "../../../wailsjs/go/window/Window";
import { useWebSocket } from 'ahooks';

function Body(props: any) {
    const { readyState, sendMessage, latestMessage, disconnect, connect } = useWebSocket(
        // 'ws://localhost:34001/ws',
        'ws://39.108.140.204:34001/ws',
    );
    useEffect(() => {
        const x = (screen.width - 336) * window.devicePixelRatio - 16
        const y = (screen.height - 200) * window.devicePixelRatio - 76
        console.log(x, y)
        window.runtime.WindowSetPosition(x, y)
    }, [])

    useEffect(() => {
        if (latestMessage?.data) {
            ShowWindow()
            setTimeout(() => {
                HideWindow()
            }, 3000)
        }
    }, [latestMessage])

    return (
        <div className='body'>
            <div className='content'>
                {latestMessage?.data}
            </div>
        </div>
    )
}

export default Body;
