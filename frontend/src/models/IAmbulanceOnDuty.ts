import { AmbulancesInterface } from "./IAmbulance";

export interface AmbulanceOnDutyInterface {
    ID: number,
    Code: string,
    OnDutyDate: Date,
    Passenger: string,
    AmbulanceID:      number,
    Ambulance:        AmbulancesInterface,
  }