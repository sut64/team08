import { AmbulancesInterface } from "./IAmbulance";
import { EmployeesInterface } from "./IEmployee";
import { ProblemsInterface } from "./IProblem";

export interface AmbulanceChecksInterface {
    ID: number, 
    DateAndTime: Date,

    AmbulanceID: number,
    Ambulance: AmbulancesInterface,

    RecorderID: number,
    Recorder: EmployeesInterface,

    ProblemID: number,
    Problem: ProblemsInterface,

    DocCode: string,
    Severity: number,
    Note: string,
}