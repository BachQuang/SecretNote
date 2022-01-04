import React, { useEffect } from 'react';
import { Link, Navigate } from 'react-router-dom';
import axios from 'axios';

const Home = () => {

    const continueWithGoogle = async () => {
        try {
            const res = await axios.get(`${process.env.REACT_APP_API_URL}/auth/google/login?redirect_uri=${process.env.REACT_APP_API_URL}/google`)
            window.location.replace(res.data.redirectURL)
        } catch (err) {

        }
    };
    if (localStorage.getItem('accessToken')) {
        return <Navigate to='/posts' />
    }
    return (
        <div className='container mt-5'>
            <button className='btn btn-danger mt-3' onClick={continueWithGoogle}>
                Login With Google
            </button>
            <br />
        </div>
    );
};

export default Home;