import React from 'react'
import ReactDOM from 'react-dom/client'
import HelloWorld from "./HelloWorld.jsx"
import './index.css'


ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <div className={"container"}>
            <div className={"row"}>
                <div className={"col"}>
                    <HelloWorld msg={"Hello, World!"}/>
                </div>
            </div>
        </div>
    </React.StrictMode>,
)
