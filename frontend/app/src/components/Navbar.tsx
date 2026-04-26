import {NavLink} from "react-router-dom"

function Navbar() {
    return (

        <nav>
            <NavLink to="/Statistics">Statistics</NavLink>
            <NavLink to="/">Home</NavLink>
        </nav>
    )
}

export default Navbar