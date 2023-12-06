import { useState } from 'react';
// import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet, RandomGrandLotto } from "../wailsjs/go/main/App";
import { Badge, List } from "antd";
type res = Array<number[]>[];

function App() {
    const [count, setCount] = useState(1);
    const [res, setRes] = useState<res>([]);
    const updateCount = (e: any) => setCount(e.target.value);

    function getDate() {
        console.log("count", count)
        RandomGrandLotto(Number(count)).then((res) => {
            if (res.length > 0) {
                setRes(res)
            }
        })

    }

    console.log("res:", res)

    return (
        <div id="App">
            <div id="input" className="input-box">
                生成条数：<input id="name" className="input" onChange={updateCount} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={getDate}>生成</button>
            </div>
            <div className='container'>
                <List
                    className="demo-loadmore-list"
                    // loading={initLoading}
                    itemLayout="horizontal"
                    // loadMore={loadMore}
                    dataSource={res}
                    renderItem={(item) => (
                        <List.Item
                            actions={[<a key="list-loadmore-edit">复制</a>]}
                        >
                            <div className='item'>
                                {
                                    item.map((item, index) => ( <Badge  color={(index > 4) ? 'blue':'red'} count={item} />))
                                }
                            </div>
                        </List.Item>
                    )}
                />
            </div>
        </div>
    )
}

export default App;
