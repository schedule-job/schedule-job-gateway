package gateway

import (
	"log"
)

func (j *Job) InsertJob(item InsertItem) (string, error) {
	id, err := j.DB.InsertJob(item.Info.Name, item.Info.Description, item.Info.Author, item.Info.Members)

	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}

	errAction := j.DB.InsertAction(id, item.Action.Name, item.Action.Payload)

	if errAction != nil {
		j.DB.DeleteJob(id)
		log.Fatalln(errAction.Error())
		return "", errAction
	}

	errTrigger := j.DB.InsertTrigger(id, item.Trigger.Name, item.Trigger.Payload)

	if errTrigger != nil {
		j.DB.DeleteAction(id)
		j.DB.DeleteJob(id)
		log.Fatalln(errTrigger.Error())
		return "", errTrigger
	}

	return id, nil
}

func (j *Job) DeleteJob(job_id string) {
	j.DB.DeleteAction(job_id)
	j.DB.DeleteTrigger(job_id)
	j.DB.DeleteJob(job_id)
}
