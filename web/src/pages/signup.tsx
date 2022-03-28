import axios from 'axios';
import { useRouter } from 'next/router'
import React, { useContext, useEffect, useState } from 'react' 
import Layout from '../components/Layouts';
import { url } from './_app';


const Signup = () =>{
    const [ username, setUsername ] = useState<string>('');
    const [ email, setEmail] = useState<string>('');
    const [ password, setPassword ] = useState<string>('');
    const router = useRouter()
    const submit = async(e: any) => {
        e.preventDefault()
        const res = await axios.post(`${url}/api/user/register`, {
            username,
            email,
            password,
            password_confirm: password
        })
        if(res.data){
            router.replace('/login')
        }
    }
    useEffect(() => {
    axios.get(`${url}/api/user/me`, {
        withCredentials: true
    })
    .then(res => {
        if(!res.data.message){
            router.replace('/')
        }
    }).catch(err => console.log(err))
    },[router])
    return (
        <Layout title='login page'>
         <div className="block p-6 bg-white">
           <form>
            <div className="mb-6 form-group">
               <label className="inline-block mb-2 text-gray-700 form-label">User Name</label>
               <input type="input" className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" id="exampleInputEmail1"
                 aria-describedby="emailHelp" placeholder="Enter email" onChange={(e) => setUsername(e.target.value)} />
               <small id="emailHelp" className="block mt-1 text-xs text-gray-600">
                  
               </small>
             </div>
             <div className="mb-6 form-group">
               <label className="inline-block mb-2 text-gray-700 form-label">Email address</label>
               <input type="email" className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" id="exampleInputEmail1"
                 aria-describedby="emailHelp" placeholder="Enter email" onChange={(e) => setEmail(e.target.value)} />
               <small id="emailHelp" className="block mt-1 text-xs text-gray-600">
                  
               </small>
             </div>
             <div className="mb-6 form-group">
               <label className="inline-block mb-2 text-gray-700 form-label">Password</label>
               <input type="password" className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" id="exampleInputPassword1"
                 placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
             </div>
             <div className="mb-6 form-group form-check">
               <input type="checkbox"
                 className="float-left w-4 h-4 mt-1 mr-2 align-top transition duration-200 bg-white bg-center bg-no-repeat bg-contain border border-gray-300 rounded-sm appearance-none cursor-pointer form-check-input checked:bg-blue-600 checked:border-blue-600 focus:outline-none"
                 id="exampleCheck1" />
               <label className="inline-block text-gray-800 form-check-label">Check me out</label>
             </div>
             <button type="submit" onClick={submit} className=" px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Submit</button>
           </form>
         </div>
        </Layout>
    );
}

export default Signup