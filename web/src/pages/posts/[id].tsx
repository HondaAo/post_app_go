import axios from 'axios';
import { LegacyCharacterEncoding } from 'crypto';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react' 
import Layout from '../../components/Layouts';
import { PostProps } from '../../interface';

const Post: React.FC = ({}) =>{
    const router = useRouter();
    const intId = typeof router.query.id === "string" ? parseInt(router.query.id) : -1;
    const [ post, setPost ] = useState<PostProps | null>(null)
    useEffect(() => {
        axios.get(`/api/post/${intId}`, {
            withCredentials: true
          })
          .then(res => {
           setPost(res.data)
          }).catch(err => console.log(err))
    },[intId, router]) 
    
    return (
        <div>
            <Layout title='post page'>
             {post ? (
            <div>
               <h2 className="font-bold">{post.title}</h2> 
               <p>{post.text}</p>
            </div>
             ):(
                <div>loading...</div>
             )}
            </Layout>
        </div>
    );
}

export default Post