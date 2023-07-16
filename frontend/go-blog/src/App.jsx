import './App.css'
import {Container,} from "@chakra-ui/react";
import {Header} from "./components/header/header.jsx";
import {RecentPosts} from "./components/recent-posts/recent-posts.jsx";
import {Route, Routes} from "react-router-dom";
import {ArticleLayout} from "./layout/article-layout.jsx";

function App() {

    return (
        <Container maxW="container.xl" justifyContent='center' alignItems='center'>


            <Routes>
                {/*<Route path="/" element={<h1>Home</h1>}/>*/}
                <Route path="/view/article/:id" element={<ArticleLayout/>} />
                {/*<Route path="/about" element={<h1>About</h1>}/>*/}
            </Routes>
        </Container>)
}

export default App
