import React, { useEffect, useState } from 'react';
import { Link, Navigate, useLocation, useNavigate} from 'react-router-dom';
import axios from 'axios';
import { Form } from 'react-bootstrap';

const CreatePost = () => {
    const location = useLocation()
    const [content, setContent] = useState('')
    const [title, setTitle] = useState('')

    const history = useNavigate();
    const submitButton = (e) =>{
        e.preventDefault()
        
    }

    const backToPage = () =>{ 
        let path = "/posts"; 
        history(path);
    }

    const submitData = async (e) =>{
        e.preventDefault()
        const config = {
            headers: {
                'Content-type': 'application/json',
                Authorization: `Bearer ${localStorage.getItem('accessToken')}`
            }
        }
        const body = {
            title: title,
            content: content
        }
        
        const data = await axios.post(
            `${process.env.REACT_APP_API_URL}/api/posts`, 
            body,
            config
        )

    }
    if (localStorage.getItem('accessToken') == null) {
        return <Navigate to='/' />
    }
    return (
        <div className='container'>
            <div>
                <Form>
                    <label>Title:</label>
                    <br />
                    <input type="text" id="title" name="title" value={title}
                        onChange={(e) => setTitle(e.target.value)}/>
                    <br />
                    
                    <label>Content:</label><br />
                    <input type="text" id="content" name="content" value={content}
                        onChange={(e) => setContent(e.target.value)}/>
                    
                    <input type="submit" onClick={(e)=>{submitData(e)}}/>
                </Form>
            </div>
            <button onClick={backToPage}>
                Back
            </button>
            <br />
        </div>
    );
};

export default CreatePost;