import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import axios from "axios";
import {Image, Stack, Text} from "@chakra-ui/react";

export const ArticleLayout = () => {
    const [article, setArticle] = useState({});
    const {id} = useParams();
    const getArticle = async () => {

        try {
            const {data} = await axios.get(`http://localhost:8080/api/article/${id}`);
            console.log("data", data);
            if (data) {
                setArticle(data);

            }

            console.log("article", article);
        } catch (e) {
            console.log("ERROR", e);

        }
    };

    useEffect(() => {
        getArticle();
    }, [id]);

    return (
        <Stack>
            {article &&
                <>
                    <Text as='h1'>{article.title}</Text>
                    <Image src={article.HeroImage} />
                    <div dangerouslySetInnerHTML={{__html: article.Content}}/>
                </>

                // dangerouslySetInnerHTML={article.Content}
            }
        </Stack>)
        ;

};
