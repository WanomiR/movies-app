import {useEffect, useState} from "react";
import {Link, useNavigate, useOutletContext} from "react-router-dom";
import apiBaseUrl from "/src/index.jsx";

const ManageCatalogue = () => {
	const [movies, setMovies] = useState([]);
	const {jwtToken} = useOutletContext()

	const navigate = useNavigate();

	const fetchMovies = async () => {
		const headers =  new Headers({ 'Content-Type': 'application/json' });
		headers.append("Authorization", "Bearer " + jwtToken);
		try {
			const res = await fetch(apiBaseUrl + `/admin/movies`, {method: "GET", headers: headers})
			const data = await res.json()
			setMovies(data)
		} catch (err) {
			console.log(err)
		}
	}

	useEffect(() => {
		if (jwtToken === "") {
			navigate("/login")
			return
		}

		fetchMovies()
	}, [jwtToken, navigate]);

	return (
		<>
			<div>
				<h2>Manage Catalogue</h2>
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
								<Link to={`/admin/movies/${movie.id}`}>
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

export default ManageCatalogue