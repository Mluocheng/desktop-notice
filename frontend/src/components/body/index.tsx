import React, { FC, memo, useCallback, useEffect, useRef } from 'react';
import parse, { DOMNode, domToReact } from 'html-react-parser';
import './index.css';

type DataType = {
    type: string;
    text: string;
    html: string;
    style: React.CSSProperties
}


const Body: FC<WailsProps> = memo(({ DataText, DataType, TextAlign, DataHtml }: WailsProps) => {
    console.log(DataHtml)
    // const data = [
    //     { type: "text", text: "最新通知，链接如下：最新通知，链接如下：最新通知，链接如下：" },
    //     { type: "row", text: "www.baidu.com" },
    //     { type: "link", text: "www.baidu.com" },
    //     { type: "img", text: "https://cdn.wwads.cn/creatives/m88Dv8ffgDW2NO9TVOfe2Ee3QYRtwORH2acMe3Id.png", style: { width: "20px", height: "20px" } }
    // ]
    // const DataHtml = `<div>我是div<button id="button" onclick="window.Hug.myGlobalFunction()">按钮</button><div onClick="window.Hug.myGlobalFunction()">点击div</div></div>`;

    function handleScript(scriptContent: string) {

        const matched = scriptContent.match(/window\.([A-Za-z0-9_.]+)\((.*)\)/);
        console.log("matched:", matched)
        if (matched) {
            const funcPath = matched[1];
            // const args = matched[2].split(',').map(arg => arg.trim());
            const firstCommaIndex = matched[2].indexOf(',');
            let args = [matched[2]];
            if (firstCommaIndex != -1 ){
                args = [
                    matched[2].slice(0, firstCommaIndex), 
                    matched[2].slice(firstCommaIndex + 1)
                ]
            }
            const funcRef = funcPath.split('.').reduce((acc: { [x: string]: any; }, part) => acc[part], window);
            console.log("args:", args)
            if (typeof funcRef === 'function') {
                // 将解析出的参数传递给函数
                return funcRef.apply(null, args);
            }
        }
    }

    // 渲染text参数
    const renderNode = useCallback(() => {
        if (DataText?.includes("[{")) {
            const data: DataType[] = JSON.parse(DataText)
            console.log(data)
            return <>

                {
                    data?.map((item, index) => {
                        return <React.Fragment key={index}>
                            {(item.type === "row") && <div />}
                            {(item.type === "html") && <div dangerouslySetInnerHTML={{ __html: '<div>我是22div</div>' }} />}
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

    // 渲染html
    const renderHtmlDom = useCallback(() => {
        if (!DataHtml) return <></>
        const JSXElement = parse(DataHtml, {
            replace: (domNode: any) => {
                if (domNode?.attribs && 'onclick' in domNode.attribs) {
                    const onClick = domNode.attribs.onclick;
                    const newOnClick = () => {
                        handleScript(onClick);
                    };
                    const props = {
                        ...domNode.attribs,
                        onClick: newOnClick,

                    };
                    delete props.onclick;
                    return React.createElement(domNode?.name, props, domToReact(domNode.children));
                }
            }
        });

        return <div>{JSXElement}</div>;
    }, [DataHtml])

    return (
        <div className='body' style={{ padding: DataHtml ? '' : '8px 8px 0' }}>
            <div className='content' style={{ textAlign: TextAlign as any || "center" }}>
                {renderHtmlDom()}
                {renderNode()}
            </div>
        </div>
    )
})

export default Body;

