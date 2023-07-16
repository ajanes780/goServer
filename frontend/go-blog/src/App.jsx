import './App.css'
import {Container,} from "@chakra-ui/react";
import {Header} from "./components/header/header.jsx";
import {Route, Routes} from "react-router-dom";
import {ArticleLayout} from "./layout/article-layout.jsx";
import {HomePageLayout} from "./layout/home-page-layout.jsx";

function App() {

    return (
        <Container maxW="container.xl" justifyContent='center' alignItems='center'>

            <Header/>
            <Routes>
                <Route path="/home" element={<HomePageLayout/>}/>
                <Route path="/view/article/:id" element={<ArticleLayout/>}/>
                {/*<Route path="/about" element={<h1>About</h1>}/>*/}
            </Routes>
        </Container>)
}

export default App
