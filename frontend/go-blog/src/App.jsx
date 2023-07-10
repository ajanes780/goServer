import './App.css'
import {Container,} from "@chakra-ui/react";
import {Header} from "./components/header/header.jsx";
import {RecentPosts} from "./components/recent-posts/recent-posts.jsx";

function App() {

  return (
      <Container maxW="container.xl">
        <Header/>
        <RecentPosts/>
      </Container>)
}

export default App
