//Import dependencies
const fs = require("fs");
const zlib = require("zlib");
const https = require("https");

//Import from local or URL
const outputFile = "node_output.txt";
const inputURL =
    "https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-04-01_anthem_index.json.gz";
const writeStream = fs.createWriteStream(outputFile);

//Use Regular expression
const newYorkRegex = /\bN(?:ew\s+Y(?:ork)?|Y)\b/gi;
const jsonObjectRegex = /\{[^{}]*\}(?=\s*(?:,|$))/g;

// const writeTempStream = fs.createWriteStream("temp.txt");
let cache = "",
    totalCount = 0;
const startTime = new Date();

const request = https.get(inputURL, (response) => {
    const decompressor = zlib.createGunzip();
    response.pipe(decompressor);
    // Event handler for receiving data
    decompressor.on("data", (chunk) => {
        // Convert the chunk to string and extract URLs
        cache += chunk.toString();
        const objects = cache.match(jsonObjectRegex);
        if (objects && objects.length) {
            const lastIndex =
                cache.lastIndexOf(objects[objects.length - 1]) + objects[objects.length - 1].length;
            cache = cache.substring(lastIndex);
            objects.forEach((stringToParse) => {
                try {
                    const { description, location } = JSON.parse(stringToParse);
                    if (description && location) {
                        if (newYorkRegex.test(description) && description.includes("PPO")) {
                            writeStream.write(++totalCount + ":::" + location + "\n\n");
                            if (totalCount % 1000 === 0) {
                                const currentTime = new Date();
                                console.log(
                                    `${totalCount / 1000}k URL abstracted\n ${
                                        (currentTime - startTime) / 1000
                                    } seconds`
                                );
                            }
                        }
                    }
                } catch (error) {
                    console.log("JSON parse error ocured");
                }
            });
        }
        if (cache.length > 10000) cache = cache.substring(cache.length / 2);
    });

    // Event handler for the end of the response
    decompressor.on("end", () => {
        console.log("URL extraction completed.");
        // Close the write stream
        writeStream.end(() => {
            // Exit the program after ensuring all pending write operations are completed
            process.exit();
        });
    });

    // Event handler for errors
    decompressor.on("error", (error) => {
        console.error("Error:", error.message);
        // Close the write stream
        writeStream.end(() => {
            // Exit the program after ensuring all pending write operations are completed
            process.exit(1); // Exit with non-zero code to indicate error
        });
    });
});

// Signal handler for program termination (Ctrl+C)
process.on("SIGINT", () => {
    console.log("Exiting...");
    // Close the write stream
    writeStream.end(() => {
        // Exit the program after ensuring all pending write operations are completed
        process.exit();
    });
});
