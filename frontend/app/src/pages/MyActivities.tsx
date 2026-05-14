import { activityPayload } from "../util/CommonTypes";
import { BASE_URL, ACTIVITY_ENDPOINT } from "../util/urls";
import { useEffect, useState } from "react";
import Activity from "../components/Activity";

function MyActivities() {
    
    // const promise: Promise<activityPayload[]> = getActivities();
    // console.log(promise);

    // TODO: try to understand why we need to use useState for error and loading


    const [activities, setActivities] = useState<activityPayload[]>([]);
    const [error, setError] = useState<string | null>(null);

    // TODO: add error and loading
    useEffect(() => {
        getActivities()
            .then(
                data => {
                    setActivities(data);
                }
            )
    }, []);


    return <>
        <div>
            <h1>My Activities</h1>
            {activities.map((activity, index) => (
                <Activity key={index} {...activity} />
            ))}
        </div>
    </>;
}

async function getActivities() : Promise<activityPayload[]> {
    const response: Response = await fetch(`${BASE_URL}/${ACTIVITY_ENDPOINT}`);
    
    const profileActivities: activityPayload[] = await response.json();
    return profileActivities;
}

export default MyActivities