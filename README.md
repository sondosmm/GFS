Google File System

This project aims to establish communication between multiple devices acting as master, slave, and client.

Each slave device has a local text file. The code opens the file to calculate the frequency of each character and creates a map accordingly. This map is then sent to the master device.

The master device collects maps from all slaves and stores them in a database. It performs a mapreduce operation to determine the occurrence of each character across all maps. The resulting information is stored in a new map, which is sent to the client device.

The client device connects to the master and receives a map containing the occurrence numbers of characters.

In simpler terms, the project works as follows:

1-The master device sends a request to the slave devices to calculate the frequency of each character in their local text files and create a map accordingly.

2-The slave devices calculate the frequency of each character and create a map. They then send this map to the master device.

3-The master device collects maps from all slave devices and stores them in a database.

4-The master device performs a mapreduce operation to determine the occurrence of each character across all maps.

5-The master device stores the resulting information in a new map and sends it to the client device.

6-The client device connects to the master and receives the map containing the occurrence numbers of characters.

This project demonstrates the power of distributed computing and how it can be used to solve complex problems.
