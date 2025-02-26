
from langchain_anthropic import ChatAnthropic
from langchain_core.messages import HumanMessage
import chromadb

model = ChatAnthropic(model="claude-3-5-sonnet-20240620")
client = chromadb.Client()
collection = client.get_collection("resume")
results = collection.query(
    query_texts=["Please provide all of the skills and experience for Cole Flenniken"]
)
docs = results["documents"]
model.invoke([HumanMessage(
    content="Please build a resume based on the following documents: " + 
    str(docs))
              ])