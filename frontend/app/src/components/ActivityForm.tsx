import React, { useState } from "react"
type formData = {
    title: string,
    description: string,
    durationHours: string,
    durationMinutes: string,
    durationSeconds: string,
    activityType: string
}

const fields: { name: keyof formData;  label: string}[] = [
    {name: "title", label: "Title"},
    {name: "description", label: "Description"},
    {name: "durationHours", label: "Duration (hours)"},
    {name: "durationMinutes", label: "Duration (minutes)"},
    {name: "durationSeconds", label: "Duration (seconds)"},
    {name: "activityType", label: "Activity type"}
]


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

    // does something with the (now written to) useState
    // TODO: retrieve data from formData -> send a POST request query with it to the backend
    const handleSubmit = (e: React.form<HTMLInputElement>)

    return (
        <div>
            {fields.map((field) => (
                <div key={field.name}>
                    <form onSubmit={handleSubmit}>
                        // TODO: 
                        <button type="submit">save</button>
                        <label >{field.label}</label>
                        <input
                            type="text"
                            name={field.name}
                            value={formData[field.name]}
                            onChange={handeInputChange}
                        />
                    </form>

                </div>
            ))}
            <p>{formData.title}</p>
        </div>
    )
}

export default ActivityForm