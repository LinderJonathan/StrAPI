import ActivityForm from '../components/ActivityForm'

function Home() {
    return (
        <div>
            <h1>Home</h1>
            <ActivityForm />
        </div>
    )
    // TODO: PUT activity. Maybe at later stage in some "activities" hub
    // TODO: GET activity. Or maybe have this on page load on some "activities" page
}

export default Home