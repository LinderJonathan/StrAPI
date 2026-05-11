import { activityPayload } from "../util/CommonTypes.tsx"

function Activity(activity: activityPayload) {

    return <>
        <div className="Activity">
            <li>{activity.title}</li>
            <li>{activity.description}</li>
            <li>{activity.durationHours}</li>
            <li>{activity.durationMinutes}</li>
            <li>{activity.durationSeconds}</li>
            <li>{activity.activityType}</li>
        </div>
    </>;
}

export default Activity