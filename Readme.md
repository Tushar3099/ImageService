# Image Service

## Description

### Objective

The project focusses on handling multiple `job` requests each containing thousands of images for processing. In processing, we will download the image and calculate the twice of perimeter and store the result at image level and then sleep for 0.4sec to simulate GPU.

A single job may take hours to minutes to be completed. The jobs will have `store` that will contain multiple `images`. If a image process is failed due to some error, then the respective store will stop processing more images. The job has differnt states - `completed`, `ongoing` and `failed`.

### Solution

To make the backend performative, we have to execute Jobs concurrently, as well as we can also execute stores in a job concurrently for better performance. These jobs will be running in background and apis can interact with the jobs to check their state.

We will have four modules to implement the image processing life cycle.

1. `controller`

2. `job`

3. `store`

4. `process`

At the highest level we will have `controller`. Controller will map all the jobs with their job_id and spawn the jobs as a go-rouitne. Second, we will have `job` module which represents a single job. It executes all the stores concurrently and stores the errors(if any). Third we have the `store`, which execute the processes sequentially and stops if encounters any error. At the lowest level we have the `process` which downloads the image, calculates the perimeter and stores the result.

## Testing

Testing for each of the module is written to check all the functionality. The tests can me improved further in future.

## Future Work

- Currently we are saving all the job information in in-memmory, which is not scalable. We can run a CRON operation on the `controller` module to save all the jobs that are either failed or completed to the database and remove from the in-memory.

## Pre-Requisite

- `Golang` installed in your local machine

- `Code Editor` for developement

## Dev Environment

- **OS** : Linux(Ubuntu 20.04 LTS)

- **Editor** : Visual Studios Code

## Getting Started

1. Run the following command to start the server:

```bash
make run
```
