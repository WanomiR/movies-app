import {Link, Outlet, useNavigate} from "react-router-dom";
import {useState} from "react";
import Alert from "./components/Alert";

const App = () => {
	const [jwtToken, setJwtToken] = useState("");
	const [alertInfo, setAlertInfo] = useState({className: "d-none", message: ""});

	const navigate = useNavigate()

	const logOut = () => {
		setJwtToken("")
		navigate("/login")
	}

	return (<div className="container">
		<div className="row">
			<div className="col">
				<h1 className={"mt-3"}>Go Watch a Movie!</h1>
			</div>
			<div className="col text-end mt-3">
				{jwtToken === "" ? <Link to={"/login"}><span className={"badge bg-success"}>Login</span></Link> :
					<a href={"#!"} onClick={logOut}><span className="badge bg-danger">Logout</span></a>}
			</div>
			<hr className="mb-3"/>
		</div>

		<div className="row">
			<div className="col-md-2">
				<nav>
					<div className="list-group">
						<Link to={"/"} className={"list-group-item list-group-item-action"}>Home</Link>
						<Link to={"/movies"} className={"list-group-item list-group-item-action"}>Movies</Link>
						<Link to={"/genres"} className={"list-group-item list-group-item-action"}>Genres</Link>
						{jwtToken !== "" && <>
							<Link to={"/admin/movie/0"} className={"list-group-item list-group-item-action"}>Add
								Movie</Link>
							<Link to={"/manage-catalogue"} className={"list-group-item list-group-item-action"}>Manage
								Catalogue</Link>
							<Link to={"/graphql"} className={"list-group-item list-group-item-action"}>GraphQL</Link>
						</>}
					</div>
				</nav>
			</div>
			<div className="col-md-10">
				<Alert message={alertInfo.message} className={alertInfo.className}/>
				<Outlet context={{
					jwtToken, setJwtToken, setAlertInfo,
				}}/>
			</div>
		</div>
	</div>)
}

export default App