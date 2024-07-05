import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import App from './App'
import './index.css'
import * as cmp from './components/components.js'

const apiBaseUrl = 'http://localhost:8888'

const router = createBrowserRouter([
	{
		path: '/',
		element: <App/>,
		errorElement: <cmp.ErrorPage/>,
		children: [
			{index: true, element: <cmp.Home/>},
			{path: "/movies", element: <cmp.Movies/>,},
			{path: "/movies/:id", element: <cmp.Movie/>,},
			{path: "/genres", element: <cmp.Genres/>,},
			{path: "/admin/movie/0", element: <cmp.EditMovie/>,},
			{path: "/manage-catalogue", element: <cmp.ManageCatalogue/>,},
			{path: "/graphql", element: <cmp.GraphQL/>,},
			{path: "/login", element: <cmp.Login/>,},
		]
	},
])

ReactDOM.createRoot(document.getElementById('root')).render(
	<React.StrictMode>
		<RouterProvider router={router}/>
	</React.StrictMode>,
)

export default apiBaseUrl