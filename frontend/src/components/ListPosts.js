import React, { useEffect, useState } from 'react';
import { Link, Navigate, useNavigate} from 'react-router-dom';
import axios from 'axios';

const ListPosts = () => {

    const [pageID, setPageID] = useState(1)
    const [pageSize, setPageSize] = useState(5)
    const [listPosts, setListPosts] = useState([])

    const history = useNavigate();
    const createNewPost = () =>{ 
        let path = "/posts/create"; 
        history(path);
    }

    const getNextPage = async () =>{
        await setPageID(pageID + 1)
    }
    const getPrevPage = async () => {
        if (pageID > 1){
            setPageID(pageID-1)
        }
    }
    const logout = async() => {
        localStorage.removeItem('accessToken')
        history('/')
    }
    useEffect(async ()=> {

        const config = {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
        };
        const details = {
            'page_id': pageID,
            'page_size': pageSize
        };
        const formBody = Object.keys(details).map(key => encodeURIComponent(key) + '=' + encodeURIComponent(details[key])).join('&');
        const res = await axios.get(`${process.env.REACT_APP_API_URL}/api/posts?${formBody}`, config)
        if (res.data.length == 0 && pageID >=1){
            setPageID(pageID-1)
        }
        else{
            setListPosts(res.data)
        }
        
    }, [pageID])
    
    if (localStorage.getItem('accessToken') == null) {
        return <Navigate to='/' />
    }
    return (
        <div className='container'>
            <button onClick={createNewPost}>
                Create new note
            </button>
            <div>
            {listPosts.map(post =>(
                <Link key={post.id}
                    to={`${post.id}`}>
                    <h2>{post.title}</h2>
                    <div>{post.content}</div>
                    <div>{post.created_at}</div>
                </Link>
            ))

            }
            </div>
            <div>
            <button onClick={getNextPage}>
                NextPage
            </button>
            <button onClick={getPrevPage}>
                PrevPage
            </button>
            <br />
            </div>
            <button onClick={logout}>Logout</button>
        </div>
    );
};

export default ListPosts;