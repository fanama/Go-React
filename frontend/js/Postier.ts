import axios from 'axios'
import 'regenerator-runtime/runtime'

class Postier{

    async sendRequest(request:string,setToken){

        axios.post('/form',{name:request}).then(res=>{
            setToken(res.data)
        })

    }

    async getRequest(request:string){

        const header = `Authorization: Bearer ${request}`

        console.log({header})

        axios.get('sql/protected',{headers:{ Authorization: `Bearer ${request}` }})
            .then(res=>{
            console.log(res)
        })

    }


}

export const postier = new Postier()