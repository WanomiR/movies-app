import Ticket from '/src/images/movie_tickets.jpg'
import {Link} from "react-router-dom";

const Home = () => {
	return (
		<>
			<div className="text-center">
				<h2>Find a movie to watch tonight!</h2>
				<hr/>
				<Link to={"/movies"}>
					<img src={Ticket} alt="movie-tickets"/>
				</Link>
			</div>
		</>
	)
}

export default Home