const Alert = props => (
	<div className={"alert " + props.className} role={"alert"}>
		{props.message}
	</div>
)

export default Alert;