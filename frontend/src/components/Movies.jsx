import {useEffect, useState} from "react";
import {Link} from "react-router-dom";

const Movies = () => {
	const [movies, setMovies] = useState([]);

	useEffect(() => {
		const fetchData = async () => {
			const headers =  new Headers({ 'Content-Type': 'application/json' });
			try {
				const res = await fetch(`http://0.0.0.0:8888/movies`, {method: "GET", headers: headers})
				const data = await res.json()
				setMovies(data)
			} catch (err) {
				console.log(err)
			}
		}
		fetchData()
	}, []);

	return (
		<>
			<div>
				<h2>Movies</h2>
				<hr/>
				<table className="table table-striped table-hover">
					<thead>
					<tr>
						<th>Movie</th>
						<th>Release Date</th>
						<th>Rating</th>
					</tr>
					</thead>
					<tbody>
					{movies.map(movie => (
						<tr key={movie.id}>
							<td>
								<Link to={`/movies/${movie.id}`}>
									{movie.title}
								</Link>
							</td>
							<td>{movie.release_date}</td>
							<td>{movie.mpaa_rating}</td>
						</tr>
					))}
					</tbody>
				</table>
			</div>
		</>
	)
}

export default Movies