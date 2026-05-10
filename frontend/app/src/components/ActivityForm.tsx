import React, { useState } from "react"
import { BASE_URL, POST_ACTIVITY_ENDPOINT } from "../util/urls";

type formData = {
    title: string,
    description: string,
    durationHours: string,
    durationMinutes: string,
    durationSeconds: string,
    activityType: string
}

type activityPayload = {
    title: string,
    description: string,
    durationHours: number,
    durationMinutes: number,
    durationSeconds: number,
    activityType: number
}
const fields: { name: keyof formData;  label: string}[] = [
    {name: "title", label: "Title"},
    {name: "description", label: "Description"},
    {name: "durationHours", label: "Duration (hours)"},
    {name: "durationMinutes", label: "Duration (minutes)"},
    {name: "durationSeconds", label: "Duration (seconds)"},
    {name: "activityType", label: "Activity type"}
]

/*
 * Function ActivityForm
 *
 * Handles user form input to be sent to the server. 
 * Function saves user input to useState, and is formatted
 * and parsed to JSON to be sent 
 * 
 * @param: none
 * @returns: none
*/
function ActivityForm() {

    const [formData, setFormData] = useState<formData>({
        title: "",
        description: "",
        durationHours: "",
        durationMinutes: "",
        durationSeconds: "",
        activityType: ""
    })

    // updates the useState
    const handeInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        
        setFormData(prev => ({...prev, [name]: value}))
    }

    // Handles useState data submission
    const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
        e.preventDefault()

        // TODO: rerun the form tommorow with now parsed type members
        const payload =
        {
            ...formData,
            durationHours: Number(formData.durationHours),
            durationMinutes: Number(formData.durationMinutes),
            durationSeconds: Number(formData.durationSeconds),
            activityType: Number(formData.activityType)
        };

        await createActivity(payload);
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                {fields.map((field) => (
                    <div key={field.name}>
                            <label >{field.label}</label>
                            <input
                                type="text"
                                name={field.name}
                                value={formData[field.name]}
                                onChange={handeInputChange}
                            />
                    </div>
                ))}
                <button type="submit">Submit</button>
            </form>
            <p>{formData.title}</p>
        </div>
    )
}

/*
 * Function createActivity
 *
 * Function creates an HTTP POST request with form data for
 * an activity
 * 
 * @param formData: Form data structured for an activity
 * @returns: Server response object
*/
function createActivity(data: activityPayload) {
    console.log(data);
    console.log(JSON.stringify(data));
    return fetch(
        `${BASE_URL}/${POST_ACTIVITY_ENDPOINT}`,
        {
            method: 'POST',
            headers:
            {
                'Content-Type': 'application/json'
            },
            // TODO: fields in 'data' are not parsed to int (duration, ...)
            body: JSON.stringify(data)
        }
    );
}

export default ActivityForm