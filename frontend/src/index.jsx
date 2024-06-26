import React from 'react'
import ReactDOM from 'react-dom/client'
import AppClass from './AppClass.jsx'
import HelloWorld from "./HelloWorld.jsx"
import 'frontend/src/index.css'


ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <AppClass/>
        <HelloWorld />
    </React.StrictMode>,
)
