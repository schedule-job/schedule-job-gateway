package gateway

func (j *Job) InsertJob(item Item) (string, error) {
	id, err := j.DB.InsertJob(item.Info.Name, item.Info.Description, item.Info.Author, item.Info.Members)

	if err != nil {
		return "", err
	}

	errAction := j.DB.InsertAction(id, item.Action.Name, item.Action.Type, item.Action.Payload)

	if errAction != nil {
		j.DB.DeleteJob(id)
		return "", errAction
	}

	errTrigger := j.DB.InsertTrigger(id, item.Trigger.Name, item.Trigger.Payload)

	if errTrigger != nil {
		j.DB.DeleteAction(id)
		j.DB.DeleteJob(id)
		return "", errTrigger
	}

	return id, nil
}

func (j *Job) DeleteJob(job_id string) {
	j.DB.DeleteAction(job_id)
	j.DB.DeleteTrigger(job_id)
	j.DB.DeleteJob(job_id)
}
