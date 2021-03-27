import * as React from 'react'
import {render} from 'react-dom'
import axios from "axios"

export default function App() {

    const [username,setUsername] = React.useState("")
    const [password,setPassword] = React.useState("")


    const send = (e)=>{
        e.preventDefault()

        axios.post("/login",{username,password}).then(res=>{
            console.log(res.data)
        }).catch(err=>{
            console.log(err)
        })


    }

    return (
        <div>
            <h1>hello react</h1>

            <form onSubmit={send}  >
                <input type="text" value={username} name="username" onChange={e=>{setUsername(e.target.value)}} />
                <input type="password" value={password} name="password" onChange={e=>{setPassword(e.target.value)}} />
                <button>log in</button>
            </form>

        </div>
    )
}

render(<App/>,document.querySelector("#app"))