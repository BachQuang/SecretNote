import React, { useEffect, useState } from 'react';
import { Link, Navigate, useLocation, useNavigate} from 'react-router-dom';
import axios from 'axios';

const Post = () => {
    const location = useLocation()
    const [postID, setPostID] = useState(null)
    const [content, setContent] = useState('')
    const [title, setTitle] = useState('')
    const [createdAt, setCreateAt] = useState('')

    const history = useNavigate();
    const backToPage = () =>{ 
        let path = "/posts"; 
        history(path);
    }
    
    useEffect(async ()=> {
        try{
            const config = {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                }
            };
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/api${location.pathname}`, config)
            console.log(res.data)
            setPostID(res.data.id)
            setContent(res.data.content)
            setTitle(res.data.title)
            setCreateAt(res.data.created_at)
        }catch(err){

        }
        
    }, [postID])
    
    if (localStorage.getItem('accessToken') == null) {
        return <Navigate to='/' />
    }
    return (
        <div className='container'>
            <div key={postID}>
                <h2>{title}</h2>
                <div>{content}</div>
                <div>{createdAt}</div>
            </div>
            <button onClick={backToPage}>
                Back
            </button>
            <br />
        </div>
    );
};

export default Post;