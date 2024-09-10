const sharp = require('sharp');
const https = require('https');
const AWS = require('aws-sdk');
const S3 = new AWS.S3();

exports.handler = async (event) => {
    const imageUrl = event.imageUrl; // The image URL passed from Go Fiber
    const targetWidth = event.width || 500; // Set your default target width

    // Fetch the image from the URL
    const imageBuffer = await fetchImage(imageUrl);

    // Resize the image using Sharp
    const resizedImageBuffer = await sharp(imageBuffer)
        .resize(targetWidth)
        .toBuffer();

    // Optionally store the resized image in S3 (or return it)
    const s3Params = {
        Bucket: 'your-s3-bucket-name',
        Key: `resized-images/${new Date().getTime()}.jpg`,
        Body: resizedImageBuffer,
        ContentType: 'image/jpeg',
    };
    
    await S3.putObject(s3Params).promise();
    
    return {
        statusCode: 200,
        body: "Image resized and saved successfully!",
    };
};

// Helper function to fetch image from a URL
const fetchImage = (url) => {
    return new Promise((resolve, reject) => {
        https.get(url, (response) => {
            let data = [];
            response.on('data', (chunk) => data.push(chunk));
            response.on('end', () => resolve(Buffer.concat(data)));
            response.on('error', (err) => reject(err));
        });
    });
};
