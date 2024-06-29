import Input from "./form/Input.jsx";
import {useState} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";

const Login = () => {

	const [credentials, setCredentials] = useState({email: "", password: ""});
	const {setJwtToken, setAlertInfo} = useOutletContext()

	const navigate = useNavigate()

	const handleSubmit = async (e) => {
		e.preventDefault()

		// build the request payload
		let payload = {
			email: credentials.email,
			password: credentials.password,
		}

		const headers =  new Headers({ 'Content-Type': 'application/json' });
		const requestOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(payload),
		}

		try {
			const res = await fetch(`http://localhost:8888/authenticate`, requestOptions)
			const data = await res.json()
			if (data.error) {
				setAlertInfo({className: "alert-danger", message: data.message})
			} else {
				setJwtToken(data.access_token)
				setAlertInfo({className: "d-none", message: ""})
				navigate("/")
			}
		} catch (e) {
			setAlertInfo({className: "alert-danger", message: e.message})
		}
	}

	return (
		<>
			<div className={"col-md-6 offset-md-3"}>
				<h2>Login</h2>
				<hr/>
				<form onSubmit={handleSubmit}>
					<Input
						title={"Email address"}
						type="email"
						className={"form-control"}
						name={"email"}
						autoComplete="email-new"
						onChange={(e) => setCredentials({...credentials, email: e.target.value})}
					/>
					<Input
						title={"Password"}
						type="password"
						className={"form-control"}
						name={"password"}
						autoComplete="password-new"
						onChange={(e) => setCredentials({...credentials, password: e.target.value})}
					/>
					<input
						type="submit"
						className={"btn btn-primary"}
						value={"Login"}
					/>
				</form>
			</div>
		</>
	)
}

export default Login