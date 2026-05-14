import { activityPayload } from "../util/CommonTypes.tsx"

function Activity(props: activityPayload) {

    return <>
        <div className="Activity">
            <li>{props.title}</li>
            <li>{props.description}</li>
            <li>{props.durationHours}</li>
            <li>{props.durationMinutes}</li>
            <li>{props.durationSeconds}</li>
            <li>{props.activityType}</li>
        </div>
    </>;
}

export default Activity