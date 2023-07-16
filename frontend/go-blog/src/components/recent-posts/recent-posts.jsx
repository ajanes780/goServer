import axios from "axios";
import {useEffect, useState} from "react";
import {SimpleGrid, Spinner, Stack, Text} from "@chakra-ui/react";
import {SimpleCard} from "../cards/simple-card.jsx";


export const RecentPosts = () => {
  const [articles, setArticles] = useState([]);

  const getArticles = async () => {
    try {
      const result = await axios.get('http://localhost:8080/api/articles')
      console.log(result.data)
      setArticles(result.data)

    } catch (e) {
      console.log("Error", e)
    }

  };
  useEffect(() => {
    getArticles();
  }, []);


  return (
      <Stack direction={'row'} justifyContent='center' mt={10}>

        {
          articles.length ?
              <SimpleGrid columns={[1, 2, 3]} spacing={10}  justifyContent='center'>
                {articles.map((article) => {
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

                })}
              </SimpleGrid> :
              <Stack>
                <Text>
                  Loading...
                </Text>
                <Spinner
                    thickness='4px'
                    speed='0.65s'
                    emptyColor='gray.200'
                    color='blue.500'
                    size='xl'
                />
              </Stack>
        }

      </Stack>
  )
};

