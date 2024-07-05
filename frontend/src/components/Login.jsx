import Input from "./form/Input.jsx";
import {useState} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";
import apiBaseUrl from "/src/index.jsx";

const Login = () => {

	const [credentials, setCredentials] = useState({email: "", password: ""});
	const {setJwtToken, setAlertInfo, toggleRefresh} = useOutletContext()

	const navigate = useNavigate()

	const handleSubmit = async e => {
		e.preventDefault()

		// build request payload
		let payload = {
			email: credentials.email,
			password: credentials.password,
		}

		const requestOptions = {
			method: "POST",
			headers: {'Content-Type': 'application/json'},
			credentials: "include",
			body: JSON.stringify(payload),
		}

		try {
			const res = await fetch(apiBaseUrl + `/authenticate`, requestOptions)
			const data = await res.json()
			if (data.error) {
				setAlertInfo({message: data.message, className: "alert-danger"})
			} else {
				setJwtToken(data.access_token)
				setAlertInfo({message: "", className: "d-none"})
				toggleRefresh(true)
				navigate("/")
			}

		} catch (error) {
			setAlertInfo({message: error.message, className: "alert-danger"})
		}
	}

	return (
		<>
			<div className={"col-md-6 offset-md-3"} >
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
						placeholder={"admin@example.com"}
					/>
					<Input
						title={"Password"}
						type="password"
						className={"form-control"}
						name={"password"}
						placeholder={"secret"}
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