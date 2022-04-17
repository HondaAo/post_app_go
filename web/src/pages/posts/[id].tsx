import axios from 'axios';
import { LegacyCharacterEncoding } from 'crypto';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react' 
import Layout from '../../components/Layouts';
import { Reply } from '../../components/Reply';
import { PostProps, ReplyProps } from '../../interface';

const Post: React.FC = ({}) =>{
    const router = useRouter();
    const intId = typeof router.query.id === "string" ? parseInt(router.query.id) : -1;
    const [ post, setPost ] = useState<PostProps | null>(null)
    const [ replies, setReplies ] = useState<ReplyProps[] | null>(null)
    useEffect(() => {
        axios.get(`/api/post/${intId}`, {
            withCredentials: true
          })
          .then(res => {
           setPost(res.data)
           setReplies(res.data.replies)
          }).catch(err => console.log(err))
    },[intId, router]) 
    
    return (
        <div>
            <Layout title='post page'>
             {post ? (
            <div>
               <h2 className="font-bold">{post.title}</h2> 
               <p className='my-4'>{post.text}</p>
               <p>
                  points: {post.points}
               </p>
                {replies ? replies.map(reply  => {
                return (
                <div key={reply.id}>
                {reply.reply_id != 0 ? (
                    <div>

                    </div>
                ): (
                     <Reply key={reply.id} reply={reply} />
                )}
                </div>
                )
                }): (
                    <div>
                     no reply
                    </div>
                )}
            </div>
             ):(
                <div>loading...</div>
             )}
            </Layout>
        </div>
    );
}

export default Post
