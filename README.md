# Master Slave Database Architecture

The objective of this project is to establish communication among multiple devices functioning as master, slave, and client.

Tools: Visual Studio Code, MySQL Workbench.

Languages: go.

The code assumes that each slave device possesses a local text file.
It opens the file to calculate the frequency of each character and creates a map accordingly.
This map is then sent to the master device.

The master device collects maps from all slaves and stores them in a database. 
It performs a mapreduce operation to determine the occurrence of each character across all maps.
The resulting information is stored in a new map, which is sent to the client device.

The client device connects to the master and receives a map containing the occurrence numbers of characters.

To execute this project successfully, it is necessary to disable the Windows Firewall and configure the IP addresses of the devices.
