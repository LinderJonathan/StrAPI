import { activityPayload } from "../util/CommonTypes";
import { BASE_URL, ACTIVITY_ENDPOINT } from "../util/urls";

function MyActivities() {
    
    const promise: Promise<activityPayload[]> = getActivities();
    console.log(promise);
    // TODO: try to understand why we need to use useEffect and useState here. Then use it

    return <>
        <div>
            <h1>My Activities</h1>
        </div>
    </>;
}

async function getActivities() : Promise<activityPayload[]> {
    const response: Response = await fetch(`${BASE_URL}/${ACTIVITY_ENDPOINT}`);
    
    const profileActivities: activityPayload[] = await response.json();
    return profileActivities;
}

export default MyActivities