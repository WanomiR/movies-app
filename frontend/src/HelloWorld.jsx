import React, {useState} from 'react';

const HelloWorld = props => {
	const [isTrue, setIsTrue] = useState(true)

	return (
		<>
			<hr/>
			<h1>{props.msg}</h1>
			<hr/>
			{isTrue
                ? <p>isTrue</p>
                : <p>isFalse</p>
			}
            <a href="#" className="btn btn-outline-secondary" onClick={() => {setIsTrue(!isTrue)}}>Toggle isTrue</a>
		</>
	)
}

export default HelloWorld