import React, { useEffect } from 'react' 
import { Post }  from '../interface' 
import Image from 'next/image'
import axios from 'axios'
import { url } from '../pages/_app'
import { useRouter } from 'next/router'
import { route } from 'next/dist/server/router'

interface PostProps {
    post: Post
}

export const PostComponent: React.FC<PostProps> = ({post}) =>{
    const router = useRouter()
    const isAuth = async() => {
        const response = await axios.get(`${url}/api/user/me`, { withCredentials: true })
        if(response.data.message == "not authed"){
           return false
        }
        return true
    }
    const upVote = async (id: number) => {
        const auth = await isAuth() 
        if(auth){
        const res = await axios.post(`${url}/api/vote`,{
            post_id: id,
            value: 1
        }, { withCredentials: true})

        if(res.data.points != post.points){
            post.points = res.data.points 
        }
        }else{
            router.replace('/login')
        }
    }
    const downVote = async (id: number) => {
        await isAuth()
        const res = await axios.post(`${url}/api/vote`, {
            post_id: id,
            value: -1
        }, { withCredentials: true})

        if(res.data.points != post.points){
            post.points = res.data.points
        }
    }
    return (
        <div className="flex w-full max-w-sm my-2 border-t border-b border-l border-gray-400 lg:max-w-full">
         <div className='w-1/12 bg-gray-50' style={{ 'textAlignLast': 'center', 'alignSelf': 'center'}}>
             <div>
                 <Image src='/up-arrow-icon.png' height="60" width="60" alt='' onClick={() => upVote(post.id)} />
             </div>
             <div>{post.points}</div>
             <div>
             <Image src='/down-arrow-icon.png' height="60" width="60" alt='' onClick={() => downVote(post.id)} />
             </div>
         </div>
         <div className="flex flex-col justify-between w-11/12 p-4 leading-normal bg-white border-b border-r border-gray-400 rounded-b lg:border-t lg:border-gray-400 lg:rounded-b-none lg:rounded-r">
           <div className="mb-8">
             <p className="flex items-center text-sm text-gray-600">
               <svg className="w-3 h-3 mr-2 text-gray-500 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                 <path d="M4 8V6a6 6 0 1 1 12 0v2h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-8c0-1.1.9-2 2-2h1zm5 6.73V17h2v-2.27a2 2 0 1 0-2 0zM7 6v2h6V6a3 3 0 0 0-6 0z" />
               </svg>
               Members only
             </p>
             <div className="mb-2 text-xl font-bold text-gray-900">{post.title}</div>
             <p className="text-base text-gray-700">{post.text}</p>
           </div>
           <div className="flex items-center">
             <div className="text-sm">
               <p className="leading-none text-gray-900">Jonathan Reinink</p>
               <p className="text-gray-600">Aug 18</p>
             </div>
           </div>
         </div>
       </div>
    );
}