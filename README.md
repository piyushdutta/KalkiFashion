# KalkiFashion

To do
(External Python script to fetch the events data)
1> Create a go routine for triggering the external python script. 
 1.a> Check for the file length of the csv is greater than 1.
 1.b> The wrtie is completed only then resolve the go thread.
 1.c> How to make the date in query.json set to be today's date in yyyy-mm-dd format.

2> Query for SKU id's
 2.a> Initial query for fetching the SKU id is done.
 2.b> Using a if - else case ask for user permission for the number of sku id's data they want (by default only 4)
 2.c> Re-query using a function to fetch the 4 items and store it in a struct.
 2.d> Pack the data as a json object.(Done)
 2.e> Wrtie the data in a CSV file
 
3> Trigger the external go script to upload the data as user property
 3.a> Create a gocrout to execute the script 
 3.b> Check for the file size is greater than 1KB.
 3.c> Initate the user property upload. 
 3.d> wait till completion is recieved from the script.
 
4> Timestamp limitiation
 4.a> Put a end date on the script using the system timestamp.
 4.b> Once the timestamp matches stop the execution of the script and show a banner to contact support@clevertap.com
 4.c> Update the CSM accordingly that not more than 1 month will be give to the client to test this as a feature.
