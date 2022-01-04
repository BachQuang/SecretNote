import React, { useEffect } from 'react';
import { Link, useLocation, Navigate, useNavigate } from 'react-router-dom';
import axios from 'axios';
import queryString from 'query-string';

const Google = () => {
    let location = useLocation();

    const history = useNavigate();
    const handlerFunc = async (state, code) => {
        const config = {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        };
        const details = {
            'state': state,
            'code': code
        };
        const formBody = Object.keys(details).map(key => encodeURIComponent(key) + '=' + encodeURIComponent(details[key])).join('&');
        try {
            const res = await axios.post(`${process.env.REACT_APP_API_URL}/auth/google/callback?${formBody}`, config);
            console.log(res)
            localStorage.setItem("accessToken", res.data.access_token)
            console.log(res.data)
            history("/posts")

        } catch (err) {

        }

    }
    useEffect(() => {
        const values = queryString.parse(location.search);
        const state = values.state ? values.state : null;
        const code = values.code ? values.code : null;


        if (state && code) {
            handlerFunc(state, code);
        }
    }, [location]);

    if (localStorage.getItem('accessToken')) {
        return <Navigate to='/posts' />
    }
    return (
        <div className='container'>
            Welcome to secretnote system
        </div>
    );
};

export default Google;