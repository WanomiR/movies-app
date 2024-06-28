import Input from "./form/Input.jsx";
import {useState} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";

const Login = () => {

	const [credentials, setCredentials] = useState({email: "", password: ""});
	const {setJwtToken, setAlertInfo} = useOutletContext()

	const navigate = useNavigate()

	const handleSubmit = (e) => {
		e.preventDefault()
		console.log("email/pass", credentials.email, credentials.password)

		if (credentials.email === "admin@example.com") {
			setJwtToken("abc")
			setAlertInfo({className: "d-none", message: ""})
			navigate("/")
		} else {
			setAlertInfo({className: "alert-danger", message: "Invalid credentials"})
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