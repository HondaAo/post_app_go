import { getByTestId } from '@testing-library/react';
import axios from 'axios';
import { GetServerSideProps } from 'next/types';
import React, { useEffect, useState } from 'react' 
import { ReplyProps } from '../interface';

interface ReplyComponentProps {
    reply: ReplyProps
}

export const Reply: React.FC<ReplyComponentProps> = ({reply}) =>{
    const [ rep, setRep ] = useState<ReplyProps[] | null>(null)
    const checkRep = async() => {
        const res = await axios.get(`/api/comment/${reply.id}`, { withCredentials: true})
        setRep(res.data)
    }
    console.log(rep)
    return (
        <div>
            {reply.text}
                <div onClick={checkRep}>
                 check reply
                </div>
            {rep && rep!.map(r => {
            <div>
                <div id="rep">{r.text}</div>
            </div> 
            })}
        </div>
    );
}