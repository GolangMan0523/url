## Languages / Tech stacks
I built both Golang project and Node.js project
./go     Go project (main.go)
./js    Node project (index.js)

## How to run
### Go
1) Make sure you have installed go lastest version.
2) cd ./go
3) go run main.go
4) the output would be go_output.txt 
### JavaScript
1) Make sure you have already installed node v>20
2) cd ./js    npm install
3) npm run start
4) the output would be node_output.txt

## description
1) HTTP Request and Decompression:
The program starts by making an HTTP GET request to retrieve the gzip-compressed JSON file from a specified URL.
It then creates a gzip decompressor to read and decompress the data stream.
Regular Expressions:

2) Regular expressions are used to identify and extract JSON objects within each line of the data stream.
One regular expression (jsonObjectRegex) is used to match and extract individual JSON objects from the data stream.
Another regular expression (newYorkRegex) is employed to identify descriptions containing mentions of New York.
Data Processing Loop:

3) The program reads the decompressed data stream line by line.
For each line, it extracts JSON objects using the defined regular expressions.
It then parses each JSON object to extract the description and location fields.

4) Conditional Filtering and Output:
If a description contains a mention of New York and includes the term "PPO," the corresponding location is written to an output file (go_output.txt).
The program also keeps track of the total number of matching URLs processed and prints progress updates every 1000 URLs.

5) Time Tracking:
The program records the start and end times to measure the total execution time.

## Inputs
The input to this takehome is the Anthem machine readable index file
(https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-04-01_anthem_index.json.gz) 

You should write code that can open the machine readable index file and extract some in-network file URLs from it according to the schema published at [CMS' transparency in coverage repository](https://github.com/CMSgov/price-transparency-guide/tree/master/schemas/table-of-contents), so you can extract the data requested.

## Outputs
Your output should be the list of machine readable file URLs corresponding to Anthem's PPO in New York state. 

## Hints and Pointers
As you start working with the index, you'll quickly notice that the index file itself is extremly large, data is very frequently repeated, plan descriptions seem to contain random businesses in various regions around the country, and that there are a handful of different url styles. 

- How do you handle the file size and format efficiently, when the uncompressed file will exceed memory limitations on most systems? 
In this case, I used buffer stream for large file manage
- When you look at your output URL list, which segments of the URL are changing, which segments are repeating, and what might that mean?

- Is the 'description' field helpful? Is it complete? Does it change relative to 'location'? Is Highmark the same as Anthem?

- Anthem has an interactive MRF lookup system. 

This lookup can be used to gather additional information - but it requires you to input the EIN or name of an employer who offers an Anthem health plan: [Anthem EIN lookup](https://www.anthem.com/machine-readable-file/search/). 
How might you find a business likely to be in the Anthem NY PPO? How can you use this tool to confirm if your answer is complete?



### Deliverable
You should [send us](mailto:engineering@serifhealth.com) a link to a public repository or zip file that contains at miminum:
1. The script or code used to parse the file and produce output. 
2. The setup or packaging file(s) required to bootstrap and execute your solution code
3. The output URL list.
4. A README file, explaining your solution, how long it took you to write, how long it took to run, and the tradeoffs you made along the way. 
