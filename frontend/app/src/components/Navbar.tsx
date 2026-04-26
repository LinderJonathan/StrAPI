import {NavLink} from "react-router-dom"

function Navbar() {
    return (

        <nav>
            <NavLink to="/Statistics">Statistics</NavLink>
            <NavLink to="/">Homea</NavLink>
        </nav>
    )
}

export default Navbar