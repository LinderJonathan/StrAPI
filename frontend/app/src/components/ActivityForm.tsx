import { useState } from "react"

function ActivityForm() {

    const fields = [
        {name: "title", label: "Title"},
        {name: "description", label: "Description"},
        {name: "durationHours", label: "Duration (hours)"},
        {name: "durationMinutes", label: "Duration (minutes)"},
        {name: "durationSeconds", label: "Duration (seconds)"},
        {name: "activityType", label: "Activity type"}
    ]

    const [formData, setFormData] = useState({
        title: "",
        description: "",
        durationHours: "",
        durrationMinutes: "",
        durationSeconds: "",
        activityType: ""
    })

    // TODO: see that formData.title is updated
    // TODO: make it general for all the fields
    return (
        <div>
            {fields.map((field) => (
                <div key={field.name}>
                    <label >{field.label}</label>

                    {field.name == "title" ? (
                        <input
                            type="text"
                            name="title"
                            value={formData.title}
                            onChange={(e) =>
                                setFormData({ ...formData, title: e.target.value})
                            }
                            />
                    ) :    
                    <input type="text" />
                    }
                </div>
            ))}
            <p>{formData.title}</p>
        </div>
    )
}

export default ActivityForm