import React, { FC, memo, useCallback, useEffect } from 'react';
import './index.css';

type DataType = {
    type: string;
    text: string;
    style: React.CSSProperties
}


const Body: FC<WailsProps> = memo(({ DataText, DataType, TextAlign }: WailsProps) => {

    const data = [
        { type: "text", text: "最新通知，链接如下：最新通知，链接如下：最新通知，链接如下：" },
        { type: "row", text: "www.baidu.com" },
        { type: "link", text: "www.baidu.com" },
        { type: "img", text: "https://cdn.wwads.cn/creatives/m88Dv8ffgDW2NO9TVOfe2Ee3QYRtwORH2acMe3Id.png", style: { width: "20px", height: "20px" } }
    ]

    const renderNode = useCallback(() => {

        if (DataText?.includes("[{")) {
            const data: DataType[] = JSON.parse(DataText)
            console.log(data)
            return <>
                {
                    data?.map((item, index) => {
                        return <React.Fragment key={index}>
                            {(item.type === "row") && <div />}
                            {(item.type === "text") && <span {...item}>{item.text}</span>}
                            {(item.type === "link") && <a href={item.text} target='blank' {...item}>{item.text}</a>}
                            {(item.type === "img") && <img className='dataImg' src={item.text} alt="图片" {...item} />}
                        </React.Fragment>
                    })
                }
            </>
        }

        return <>
            {(DataType === "text") && DataText}
            {(DataType === "link") && <a href={DataText} target='blank'>{DataText}</a>}
            {(DataType === "img") && <img className='dataImg' src={DataText} alt="图片" />}
        </>
    }, [DataText])



    return (
        <div className='body'>
            <div className='content' style={{ textAlign: TextAlign as any || "center" }}>
                {renderNode()}
            </div>
        </div>
    )
})

export default Body;

