import axios from "axios";
import {useEffect, useState} from "react";
import {SimpleGrid} from "@chakra-ui/react";
import {SimpleCard} from "../cards/simple-card.jsx";


export const RecentPosts =  () => {
  const [articles, setArticles] = useState([]);

  const getArticles = async () => {
    try{
    const result = await axios.get('http://localhost:8080/api/articles')
    console.log(result.data)
    setArticles(result.data)

    }catch (e) {
        console.log("Error", e)
    }

  };
  useEffect(() => {
    getArticles();
  }, []);


  return (
      <SimpleGrid columns={2} spacing={10} mt={10}>

        {articles.length  ? articles.map((article) => {
          return (
              <SimpleCard
                  key={article.id}
                  Title={article.Title}
                  Summary={article.Summary}
                  AuthorName={article.AuthorName}
                  WrittenOn={article.WrittenOn}
                  HeroImage={article.HeroImage}
              />
          );
        }) : <div>loading...</div>}
      </SimpleGrid>
  );

};

