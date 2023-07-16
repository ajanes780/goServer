import {Stack} from "@chakra-ui/react";
import {RecentPosts} from "../components/recent-posts/recent-posts.jsx";

export const  HomePageLayout = (props) => {
    return (
       <Stack>
           <RecentPosts/>
       </Stack>
    );
}

