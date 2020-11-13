import * as React from 'react'
import { useState } from 'react'
import {render} from 'react-dom'
import { postier } from '../Postier'


export default function App() {

    const [user, setUser] = useState<string>("")
    const [token, setToken] = useState<string>("")

    const [inputtoken, setInputToken] = useState<string>("")



    const getToken = (e:React.FormEvent<HTMLFormElement>)=>{
        e.preventDefault()
        postier.sendRequest(user,setToken)
    }

    return (
        <div>
            <form onSubmit={getToken}>
                <h2>username</h2>
                <input value={user} onChange={(e)=>{setUser(e.target.value)}} />
                <button>show all</button>
            </form>
            <form>
                {(token!="")?`you are ${token}`:null}
            </form>
        </div>
    )
}

render(<App/>,document.querySelector("#app"))