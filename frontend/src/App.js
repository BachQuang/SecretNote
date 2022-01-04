import logo from './logo.svg';
import './App.css';
import  { Container} from 'react-bootstrap';

import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home'
import Google from './components/Google';
import ListPosts from './components/ListPosts';
import Post from './components/Post';
import CreatePost from './components/CreatePost';

const App = () => (
  <div>
    <Router>
      <Routes>
        <Route exact path="/" element={<Home/>} />
        <Route exact path="/google" element={<Google/>} />
        <Route exact path='/posts' element={<ListPosts/>} />
        <Route exact path='/posts/:id' element={<Post/>} />
        <Route exact path='/posts/create' element={<CreatePost/>} />
      </Routes>
    </Router>
  </div>
);

export default App;
